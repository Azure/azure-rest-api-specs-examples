import com.azure.core.util.Context;

/** Samples for SecureScores List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScores/ListSecureScores_example.json
     */
    /**
     * Sample code: List secure scores.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listSecureScores(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.secureScores().list(Context.NONE);
    }
}
