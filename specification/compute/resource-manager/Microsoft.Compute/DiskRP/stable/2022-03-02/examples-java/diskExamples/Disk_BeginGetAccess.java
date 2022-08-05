import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.AccessLevel;
import com.azure.resourcemanager.compute.models.GrantAccessData;

/** Samples for Disks GrantAccess. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-02/DiskRP/examples/diskExamples/Disk_BeginGetAccess.json
     */
    /**
     * Sample code: Get a sas on a managed disk.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASasOnAManagedDisk(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDisks()
            .grantAccess(
                "myResourceGroup",
                "myDisk",
                new GrantAccessData().withAccess(AccessLevel.READ).withDurationInSeconds(300),
                Context.NONE);
    }
}
