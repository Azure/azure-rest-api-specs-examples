
/**
 * Samples for SshPublicKeys GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/sshPublicKeyExamples/SshPublicKey_Get.json
     */
    /**
     * Sample code: Get an ssh public key.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAnSshPublicKey(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSshPublicKeys().getByResourceGroupWithResponse("myResourceGroup",
            "mySshPublicKeyName", com.azure.core.util.Context.NONE);
    }
}
