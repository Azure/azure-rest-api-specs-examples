import com.azure.resourcemanager.logic.models.KeyType;
import com.azure.resourcemanager.logic.models.RegenerateActionParameter;

/** Samples for Workflows RegenerateAccessKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_RegenerateAccessKey.json
     */
    /**
     * Sample code: Regenerate the callback URL access key for request triggers.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void regenerateTheCallbackURLAccessKeyForRequestTriggers(
        com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflows()
            .regenerateAccessKeyWithResponse(
                "testResourceGroup",
                "testWorkflowName",
                new RegenerateActionParameter().withKeyType(KeyType.PRIMARY),
                com.azure.core.util.Context.NONE);
    }
}
