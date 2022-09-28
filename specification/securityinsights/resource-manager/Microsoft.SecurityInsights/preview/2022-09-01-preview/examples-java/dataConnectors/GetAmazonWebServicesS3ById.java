import com.azure.core.util.Context;

/** Samples for DataConnectors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/GetAmazonWebServicesS3ById.json
     */
    /**
     * Sample code: Get an Aws S3 data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAnAwsS3DataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .getWithResponse("myRg", "myWorkspace", "afef3743-0c88-469c-84ff-ca2e87dc1e48", Context.NONE);
    }
}
