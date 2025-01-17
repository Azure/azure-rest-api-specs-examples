
import com.azure.resourcemanager.mixedreality.models.Identity;
import com.azure.resourcemanager.mixedreality.models.RemoteRenderingAccount;
import com.azure.resourcemanager.mixedreality.models.ResourceIdentityType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for RemoteRenderingAccounts Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/remote-rendering/
     * Patch.json
     */
    /**
     * Sample code: Update remote rendering account.
     * 
     * @param manager Entry point to MixedRealityManager.
     */
    public static void
        updateRemoteRenderingAccount(com.azure.resourcemanager.mixedreality.MixedRealityManager manager) {
        RemoteRenderingAccount resource = manager.remoteRenderingAccounts()
            .getByResourceGroupWithResponse("MyResourceGroup", "MyAccount", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("hero", "romeo", "heroine", "juliet"))
            .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED)).apply();
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
