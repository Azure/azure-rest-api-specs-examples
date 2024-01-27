
/** Samples for CustomDomains Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/CustomDomains_Get.json
     */
    /**
     * Sample code: CustomDomains_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customDomainsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getCustomDomains().getWithResponse("RG", "profile1", "endpoint1",
            "www-someDomain-net", com.azure.core.util.Context.NONE);
    }
}
