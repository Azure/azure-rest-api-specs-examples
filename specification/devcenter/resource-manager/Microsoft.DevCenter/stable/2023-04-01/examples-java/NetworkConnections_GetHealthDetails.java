/** Samples for NetworkConnections GetHealthDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/NetworkConnections_GetHealthDetails.json
     */
    /**
     * Sample code: NetworkConnections_GetHealthDetails.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void networkConnectionsGetHealthDetails(
        com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .networkConnections()
            .getHealthDetailsWithResponse("rg1", "eastusnetwork", com.azure.core.util.Context.NONE);
    }
}
