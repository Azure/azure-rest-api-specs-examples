
/**
 * Samples for Connector Validate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * ValidateConnectorSuccess.json
     */
    /**
     * Sample code: ValidateConnectorSuccess.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void validateConnectorSuccess(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.connectors().validate("00000000-0000-0000-0000-000000000000", "test-rg", "westus", "connectorName",
            com.azure.core.util.Context.NONE);
    }
}
