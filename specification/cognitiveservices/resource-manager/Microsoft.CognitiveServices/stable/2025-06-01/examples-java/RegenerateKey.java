
import com.azure.resourcemanager.cognitiveservices.models.KeyName;
import com.azure.resourcemanager.cognitiveservices.models.RegenerateKeyParameters;

/**
 * Samples for Accounts RegenerateKey.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
     * RegenerateKey.json
     */
    /**
     * Sample code: Regenerate Keys.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void regenerateKeys(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().regenerateKeyWithResponse("myResourceGroup", "myAccount",
            new RegenerateKeyParameters().withKeyName(KeyName.KEY2), com.azure.core.util.Context.NONE);
    }
}
