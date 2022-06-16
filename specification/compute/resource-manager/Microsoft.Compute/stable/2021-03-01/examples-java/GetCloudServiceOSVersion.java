import com.azure.core.util.Context;

/** Samples for CloudServiceOperatingSystems GetOSVersion. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/GetCloudServiceOSVersion.json
     */
    /**
     * Sample code: Get Cloud Service OS Version.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCloudServiceOSVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServiceOperatingSystems()
            .getOSVersionWithResponse("westus2", "WA-GUEST-OS-3.90_202010-02", Context.NONE);
    }
}
