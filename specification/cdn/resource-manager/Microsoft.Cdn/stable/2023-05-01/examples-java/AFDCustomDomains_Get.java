
/** Samples for AfdCustomDomains Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/AFDCustomDomains_Get.json
     */
    /**
     * Sample code: AFDCustomDomains_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDCustomDomainsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getAfdCustomDomains().getWithResponse("RG", "profile1", "domain1",
            com.azure.core.util.Context.NONE);
    }
}
