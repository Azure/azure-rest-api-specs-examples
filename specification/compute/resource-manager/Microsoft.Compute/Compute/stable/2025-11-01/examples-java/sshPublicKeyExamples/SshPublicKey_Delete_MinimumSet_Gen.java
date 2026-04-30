
/**
 * Samples for SshPublicKeys Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/sshPublicKeyExamples/SshPublicKey_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: SshPublicKey_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void sshPublicKeyDeleteMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSshPublicKeys().deleteWithResponse("rgcompute", "aaaaaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
