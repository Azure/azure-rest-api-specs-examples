
/**
 * Samples for ApiCollections ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2023-11-15/examples/ApiCollections/
     * APICollections_ListByResourceGroup_example.json
     */
    /**
     * Sample code: Gets a list of API collections within a resource group that have been onboarded to Microsoft
     * Defender for APIs.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getsAListOfAPICollectionsWithinAResourceGroupThatHaveBeenOnboardedToMicrosoftDefenderForAPIs(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.apiCollections().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
