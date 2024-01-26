
/**
 * Samples for CloudServiceOperatingSystems GetOSVersion.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudServiceOSVersion_Get.json
     */
    /**
     * Sample code: Get Cloud Service OS Version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCloudServiceOSVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServiceOperatingSystems()
            .getOSVersionWithResponse("westus2", "WA-GUEST-OS-3.90_202010-02", com.azure.core.util.Context.NONE);
    }
}
