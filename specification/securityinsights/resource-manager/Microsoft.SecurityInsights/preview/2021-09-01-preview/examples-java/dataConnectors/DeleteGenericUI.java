import com.azure.core.util.Context;

/** Samples for DataConnectors Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/DeleteGenericUI.json
     */
    /**
     * Sample code: Delete a GenericUI data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAGenericUIDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .deleteWithResponse("myRg", "myWorkspace", "316ec55e-7138-4d63-ab18-90c8a60fd1c8", Context.NONE);
    }
}
