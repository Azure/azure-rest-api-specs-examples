import com.azure.core.util.Context;

/** Samples for AllowedConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AllowedConnections/GetAllowedConnectionsSubscription_example.json
     */
    /**
     * Sample code: Get allowed connections on a subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getAllowedConnectionsOnASubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.allowedConnections().list(Context.NONE);
    }
}
