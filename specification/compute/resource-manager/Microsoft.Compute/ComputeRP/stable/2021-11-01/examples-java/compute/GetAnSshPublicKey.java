import com.azure.core.util.Context;

/** Samples for SshPublicKeys GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetAnSshPublicKey.json
     */
    /**
     * Sample code: Get an ssh public key.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAnSshPublicKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSshPublicKeys()
            .getByResourceGroupWithResponse("myResourceGroup", "mySshPublicKeyName", Context.NONE);
    }
}
