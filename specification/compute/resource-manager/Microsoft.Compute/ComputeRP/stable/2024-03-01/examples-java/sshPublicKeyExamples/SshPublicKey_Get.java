
/**
 * Samples for SshPublicKeys GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * sshPublicKeyExamples/SshPublicKey_Get.json
     */
    /**
     * Sample code: Get an ssh public key.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAnSshPublicKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSshPublicKeys()
            .getByResourceGroupWithResponse("myResourceGroup", "mySshPublicKeyName", com.azure.core.util.Context.NONE);
    }
}
