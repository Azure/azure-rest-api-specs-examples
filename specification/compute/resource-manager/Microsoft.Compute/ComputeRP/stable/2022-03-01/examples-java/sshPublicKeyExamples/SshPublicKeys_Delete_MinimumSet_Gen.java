import com.azure.core.util.Context;

/** Samples for SshPublicKeys Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/sshPublicKeyExamples/SshPublicKeys_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: SshPublicKeys_Delete_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sshPublicKeysDeleteMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSshPublicKeys()
            .deleteWithResponse("rgcompute", "aaaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
