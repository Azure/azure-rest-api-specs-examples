import com.azure.resourcemanager.mixedreality.models.AccountKeyRegenerateRequest;
import com.azure.resourcemanager.mixedreality.models.Serial;

/** Samples for RemoteRenderingAccounts RegenerateKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/remote-rendering/RegenerateKey.json
     */
    /**
     * Sample code: Regenerate remote rendering account keys.
     *
     * @param manager Entry point to MixedRealityManager.
     */
    public static void regenerateRemoteRenderingAccountKeys(
        com.azure.resourcemanager.mixedreality.MixedRealityManager manager) {
        manager
            .remoteRenderingAccounts()
            .regenerateKeysWithResponse(
                "MyResourceGroup",
                "MyAccount",
                new AccountKeyRegenerateRequest().withSerial(Serial.ONE),
                com.azure.core.util.Context.NONE);
    }
}
