
/**
 * Samples for DataConnectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/GetAzureActiveDirectoryById.json
     */
    /**
     * Sample code: Get an AADIP (Azure Active Directory Identity Protection) data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAnAADIPAzureActiveDirectoryIdentityProtectionDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().getWithResponse("myRg", "myWorkspace", "f0cd27d2-5f03-4c06-ba31-d2dc82dcb51d",
            com.azure.core.util.Context.NONE);
    }
}
