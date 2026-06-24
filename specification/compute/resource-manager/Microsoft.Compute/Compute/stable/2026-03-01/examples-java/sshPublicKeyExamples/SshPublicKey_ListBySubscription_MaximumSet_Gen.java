
/**
 * Samples for SshPublicKeys List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/sshPublicKeyExamples/SshPublicKey_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: SshPublicKey_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        sshPublicKeyListBySubscriptionMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSshPublicKeys().list(com.azure.core.util.Context.NONE);
    }
}
