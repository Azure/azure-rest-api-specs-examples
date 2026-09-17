
/**
 * Samples for DataConnectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/GetAmazonWebServicesS3ById.json
     */
    /**
     * Sample code: Get an Aws S3 data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAnAwsS3DataConnector(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().getWithResponse("myRg", "myWorkspace", "afef3743-0c88-469c-84ff-ca2e87dc1e48",
            com.azure.core.util.Context.NONE);
    }
}
