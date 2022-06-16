import com.azure.core.util.Context;

/** Samples for DataConnectors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/GetAmazonWebServicesCloudTrailById.json
     */
    /**
     * Sample code: Get an AwsCloudTrail data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAnAwsCloudTrailDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .getWithResponse("myRg", "myWorkspace", "c345bf40-8509-4ed2-b947-50cb773aaf04", Context.NONE);
    }
}
