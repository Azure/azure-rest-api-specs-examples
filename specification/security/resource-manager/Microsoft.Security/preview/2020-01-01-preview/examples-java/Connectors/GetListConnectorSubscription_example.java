import com.azure.core.util.Context;

/** Samples for Connectors List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/GetListConnectorSubscription_example.json
     */
    /**
     * Sample code: Get all cloud accounts connectors of a subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getAllCloudAccountsConnectorsOfASubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.connectors().list(Context.NONE);
    }
}
