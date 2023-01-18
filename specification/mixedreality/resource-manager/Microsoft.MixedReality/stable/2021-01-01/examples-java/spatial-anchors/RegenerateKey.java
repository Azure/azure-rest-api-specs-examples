import com.azure.resourcemanager.mixedreality.models.AccountKeyRegenerateRequest;
import com.azure.resourcemanager.mixedreality.models.Serial;

/** Samples for SpatialAnchorsAccounts RegenerateKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/spatial-anchors/RegenerateKey.json
     */
    /**
     * Sample code: Regenerate spatial anchors account keys.
     *
     * @param manager Entry point to MixedRealityManager.
     */
    public static void regenerateSpatialAnchorsAccountKeys(
        com.azure.resourcemanager.mixedreality.MixedRealityManager manager) {
        manager
            .spatialAnchorsAccounts()
            .regenerateKeysWithResponse(
                "MyResourceGroup",
                "MyAccount",
                new AccountKeyRegenerateRequest().withSerial(Serial.ONE),
                com.azure.core.util.Context.NONE);
    }
}
