package hell.microwars;

import com.wizzardo.http.framework.Controller;
import com.wizzardo.http.framework.template.Renderer;

import java.io.IOException;


public class DaController extends Controller {
    private static final String STATUS_STRING = "{\"status\": \"ok\"}";
    private static final String STATUS_FAILED_STRING = "{\"status\": \"not connected\"}";

    DaMagicWand daMagicWand;
    DaAttackWand daAttackWand;

    public Renderer status() {
        return renderJson(STATUS_STRING);
    }

    public Renderer test() throws IOException, InterruptedException {
        if (daAttackWand.connectionOk()) {
            return renderJson(STATUS_STRING);
        }
        return renderJson(STATUS_FAILED_STRING);
    }

    public Renderer combat() {
        daAttackWand.jab();
        daAttackWand.hook();

        return renderJson(STATUS_STRING);
    }


    public Renderer jab() {
        daAttackWand.jab();
        daAttackWand.jab();

        return renderJson(daMagicWand.getAttackResult(2));
    }

    public Renderer cross() {
        daAttackWand.jab();
        daAttackWand.jab();
        daAttackWand.cross();

        return renderJson(daMagicWand.getAttackResult(4));
    }

    public Renderer hook() {
        daAttackWand.hook();
        daAttackWand.hook();
        daAttackWand.uppercut();

        return renderJson(daMagicWand.getAttackResult(8));
    }

    public Renderer uppercut() {
        daAttackWand.cross();
        daAttackWand.hook();
        daAttackWand.uppercut();

        return renderJson(daMagicWand.getAttackResult(16));
    }

    public Renderer warmup() {
        return renderJson(daMagicWand.getAttackResult(16));
    }
}
