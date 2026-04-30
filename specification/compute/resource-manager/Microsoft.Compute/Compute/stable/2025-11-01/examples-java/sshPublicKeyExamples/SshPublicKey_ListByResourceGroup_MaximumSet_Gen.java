
/**
 * Samples for SshPublicKeys ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/sshPublicKeyExamples/SshPublicKey_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: SshPublicKey_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        sshPublicKeyListByResourceGroupMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSshPublicKeys().listByResourceGroup("rgcompute", com.azure.core.util.Context.NONE);
    }
}
