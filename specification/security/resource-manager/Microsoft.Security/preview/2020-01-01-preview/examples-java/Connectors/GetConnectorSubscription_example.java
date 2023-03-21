/** Samples for Connectors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/GetConnectorSubscription_example.json
     */
    /**
     * Sample code: Details of a specific cloud account connector.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void detailsOfASpecificCloudAccountConnector(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.connectors().getWithResponse("aws_dev1", com.azure.core.util.Context.NONE);
    }
}
