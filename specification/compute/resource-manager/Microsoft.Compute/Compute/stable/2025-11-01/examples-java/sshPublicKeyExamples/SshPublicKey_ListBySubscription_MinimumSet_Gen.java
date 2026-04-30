
/**
 * Samples for SshPublicKeys List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/sshPublicKeyExamples/SshPublicKey_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: SshPublicKey_ListBySubscription_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        sshPublicKeyListBySubscriptionMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSshPublicKeys().list(com.azure.core.util.Context.NONE);
    }
}
