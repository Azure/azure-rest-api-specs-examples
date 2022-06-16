import com.azure.core.util.Context;

/** Samples for DataConnectors Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/dataConnectors/DeleteOffice365ProjectDataConnetor.json
     */
    /**
     * Sample code: Delete an Office365 Project data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAnOffice365ProjectDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .deleteWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", Context.NONE);
    }
}
