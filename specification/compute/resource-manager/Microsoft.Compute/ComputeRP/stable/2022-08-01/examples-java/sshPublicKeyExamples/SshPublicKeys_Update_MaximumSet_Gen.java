import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.SshPublicKeyUpdateResource;
import java.util.HashMap;
import java.util.Map;

/** Samples for SshPublicKeys Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/sshPublicKeyExamples/SshPublicKeys_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: SshPublicKeys_Update_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sshPublicKeysUpdateMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSshPublicKeys()
            .updateWithResponse(
                "rgcompute",
                "aaaaaaaaaaaa",
                new SshPublicKeyUpdateResource().withTags(mapOf("key2854", "a")).withPublicKey("{ssh-rsa public key}"),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
