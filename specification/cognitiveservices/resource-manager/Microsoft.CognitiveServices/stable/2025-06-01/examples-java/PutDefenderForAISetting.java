
import com.azure.resourcemanager.cognitiveservices.models.DefenderForAISettingState;

/**
 * Samples for DefenderForAISettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
     * PutDefenderForAISetting.json
     */
    /**
     * Sample code: PutDefenderForAISetting.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        putDefenderForAISetting(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.defenderForAISettings().define("Default").withExistingAccount("resourceGroupName", "accountName")
            .withState(DefenderForAISettingState.ENABLED).create();
    }
}
