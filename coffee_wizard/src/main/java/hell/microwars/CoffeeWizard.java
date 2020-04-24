package hell.microwars;

import com.wizzardo.http.framework.WebApplication;

public class CoffeeWizard {
    public static void main(final String[] args) {
        final WebApplication application = new WebApplication(args);
        application.onSetup(app -> {
            app.setDebugOutput(false);
            app.setWorkersCount(4);
            app.setIoThreadsCount(8);
            app.getUrlMapping()
                    .append("/status", DaController.class, "status")
                    .append("/test", DaController.class, "test")
                    .append("/combat", DaController.class, "combat")
                    .append("/jab", DaController.class, "jab")
                    .append("/cross", DaController.class, "cross")
                    .append("/hook", DaController.class, "hook")
                    .append("/uppercut", DaController.class, "uppercut")
                    .append("/warmup", DaController.class, "warmup")
            ;
        });
        application.start();
    }
}
