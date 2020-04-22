package naughty.services;

import naughty.dtos.AttackResult;

import javax.inject.Singleton;
import java.util.UUID;

@Singleton
public class DaDefenceService {
    public int fibonacci(final int total) {
        if (total <= 1) return total;
        return fibonacci(total - 1) + fibonacci(total - 2);
    }

    public String newUUID() {
        return UUID.randomUUID().toString();
    }

    public AttackResult getAttackResult(final int fibNumber) {
        return new AttackResult(newUUID(), fibonacci(fibNumber));
    }
}
