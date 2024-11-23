
import com.azure.resourcemanager.cognitiveservices.models.DefenderForAISetting;
import com.azure.resourcemanager.cognitiveservices.models.DefenderForAISettingState;

/**
 * Samples for DefenderForAISettings Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/
     * UpdateDefenderForAISetting.json
     */
    /**
     * Sample code: UpdateDefenderForAISetting.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        updateDefenderForAISetting(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        DefenderForAISetting resource = manager.defenderForAISettings()
            .getWithResponse("resourceGroupName", "accountName", "Default", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withState(DefenderForAISettingState.ENABLED).apply();
    }
}
