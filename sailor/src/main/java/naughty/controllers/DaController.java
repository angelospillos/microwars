package naughty.controllers;

import io.micronaut.http.annotation.Controller;
import io.micronaut.http.annotation.Get;
import naughty.dtos.AttackResult;
import naughty.dtos.StatusResponse;
import naughty.services.DaAttackService;
import naughty.services.DaDefenceService;

@Controller("/")
public class DaController {
    private final DaDefenceService daDefenceService;
    private final DaAttackService daAttackService;
    private final StatusResponse defaultResponse = new StatusResponse();

    public DaController(final DaDefenceService daDefenceService,
                        final DaAttackService daAttackService) {
        this.daDefenceService = daDefenceService;
        this.daAttackService = daAttackService;
    }

    @Get("/status")
    public StatusResponse status() {
        return defaultResponse;
    }

    @Get("/test")
    public StatusResponse test() {
        if (daAttackService.status()) return defaultResponse;
        return new StatusResponse("not connected");
    }

    @Get("/combat")
    public StatusResponse combat() {
        daAttackService.jab();
        daAttackService.uppercut();
        return defaultResponse;
    }

    @Get("/jab")
    public AttackResult jab() {
        final AttackResult attackResult = daDefenceService.getAttackResult(2);
        daAttackService.jab();
        daAttackService.jab();
        return attackResult;
    }

    @Get("/cross")
    public AttackResult cross() {
        final AttackResult attackResult = daDefenceService.getAttackResult(4);
        daAttackService.jab();
        daAttackService.jab();
        daAttackService.cross();
        return attackResult;

    }

    @Get("/hook")
    public AttackResult hook() {
        final AttackResult attackResult = daDefenceService.getAttackResult(8);
        daAttackService.hook();
        daAttackService.hook();
        daAttackService.uppercut();
        return attackResult;

    }

    @Get("/uppercut")
    public AttackResult uppercut() {
        final AttackResult attackResult = daDefenceService.getAttackResult(16);
        daAttackService.cross();
        daAttackService.hook();
        daAttackService.uppercut();
        return attackResult;

    }
}
