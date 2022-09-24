import com.azure.core.util.Context;

/** Samples for Connectors Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/DeleteConnectorSubscription_example.json
     */
    /**
     * Sample code: Delete a cloud account connector from a subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteACloudAccountConnectorFromASubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.connectors().deleteWithResponse("aws_dev1", Context.NONE);
    }
}
