import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.AccessLevel;
import com.azure.resourcemanager.compute.models.GrantAccessData;

/** Samples for Snapshots GrantAccess. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/BeginGetAccessSnapshot.json
     */
    /**
     * Sample code: Get a sas on a snapshot.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASasOnASnapshot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSnapshots()
            .grantAccess(
                "myResourceGroup",
                "mySnapshot",
                new GrantAccessData().withAccess(AccessLevel.READ).withDurationInSeconds(300),
                Context.NONE);
    }
}
