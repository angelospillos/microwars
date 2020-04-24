package hell.microwars;

import com.wizzardo.http.framework.di.Service;

import java.util.UUID;

public class DaMagicWand implements Service {
    public int fibonacci(final int total) {
        if (total <= 1) return total;
        return fibonacci(total - 1) + fibonacci(total - 2);
    }

    public String newUUID() {
        return UUID.randomUUID().toString();
    }

    public String getAttackResult(final int fibNumber) {
        return "{ \"uuid\": \"" + newUUID() + "\", \"fib\": " + fibonacci(fibNumber) + "}";
    }
}
