package naughty.services;

import io.micronaut.http.HttpRequest;
import io.micronaut.http.HttpStatus;
import io.micronaut.http.client.RxHttpClient;
import io.micronaut.http.client.annotation.Client;
import io.micronaut.http.uri.UriBuilder;
import naughty.dtos.WinResult;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.inject.Singleton;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;

@Singleton
public class DaAttackService {
    private static final Logger log = LoggerFactory.getLogger(DaAttackService.class);

    private final RxHttpClient opponentClient;
    private final RxHttpClient refereeClient;

    private final HttpRequest<?> statusReq = HttpRequest.GET(UriBuilder.of("/status").build());
    private final HttpRequest<?> jabReq = HttpRequest.GET(UriBuilder.of("/jab").build());
    private final HttpRequest<?> crossReq = HttpRequest.GET(UriBuilder.of("/cross").build());
    private final HttpRequest<?> hookReq = HttpRequest.GET(UriBuilder.of("/hook").build());
    private final HttpRequest<?> uppercutReq = HttpRequest.GET(UriBuilder.of("/uppercut").build());

    private final HttpRequest<?> winReq = HttpRequest.POST(UriBuilder.of("/won").build(), new WinResult());

    public DaAttackService(@Client("opponent") final RxHttpClient opponentClient,
                           @Client("referee") final RxHttpClient refereeClient) {
        this.opponentClient = opponentClient;
        this.refereeClient = refereeClient;
    }

    public boolean status() {
        return HttpStatus.OK.equals(opponentClient.toBlocking().exchange(statusReq).status());
    }

    private void attack(final HttpRequest<?> request) {
        opponentClient
                .retrieve(request)
                .timeout(4, TimeUnit.SECONDS)
                .doOnError((error) -> {
                    if (error instanceof TimeoutException) {
                        log.error("TIMEOUT! WON", error);
                        refereeClient.toBlocking().retrieve(winReq);
                    } else {
                        log.error("dunno...", error);
                    }
                })
                .toObservable()
                .firstElement()
                .subscribe(x -> {
                    // Do nothing
                });

    }

    public void jab() {
        attack(jabReq);
    }

    public void cross() {
        attack(crossReq);
    }

    public void hook() {
        attack(hookReq);
    }

    public void uppercut() {
        attack(uppercutReq);
    }

}
