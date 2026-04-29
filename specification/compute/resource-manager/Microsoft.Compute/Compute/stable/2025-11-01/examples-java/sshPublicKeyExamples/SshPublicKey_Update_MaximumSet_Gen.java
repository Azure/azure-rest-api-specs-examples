
import com.azure.resourcemanager.compute.models.SshPublicKeyUpdateResource;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SshPublicKeys Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/sshPublicKeyExamples/SshPublicKey_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: SshPublicKey_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void sshPublicKeyUpdateMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSshPublicKeys().updateWithResponse(
            "rgcompute", "aaaaaaaaaaaa", new SshPublicKeyUpdateResource()
                .withTags(mapOf("key2854", "fakeTokenPlaceholder")).withPublicKey("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
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
