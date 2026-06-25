
/**
 * Samples for SshPublicKeys ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/sshPublicKeyExamples/SshPublicKey_ListByResourceGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: SshPublicKey_ListByResourceGroup_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        sshPublicKeyListByResourceGroupMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSshPublicKeys().listByResourceGroup("rgcompute", com.azure.core.util.Context.NONE);
    }
}
