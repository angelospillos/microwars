package naughty.dtos;

public class AttackResult {
    private String uuid;
    private int fib;

    public AttackResult(final String uuid, final int fib) {
        this.uuid = uuid;
        this.fib = fib;
    }

    public String getUuid() {
        return uuid;
    }

    public void setUuid(final String uuid) {
        this.uuid = uuid;
    }

    public int getFib() {
        return fib;
    }

    public void setFib(final int fib) {
        this.fib = fib;
    }
}
