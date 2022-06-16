import com.azure.core.util.Context;

/** Samples for CloudServiceOperatingSystems GetOSFamily. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/GetCloudServiceOSFamily.json
     */
    /**
     * Sample code: Get Cloud Service OS Family.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCloudServiceOSFamily(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServiceOperatingSystems()
            .getOSFamilyWithResponse("westus2", "3", Context.NONE);
    }
}
