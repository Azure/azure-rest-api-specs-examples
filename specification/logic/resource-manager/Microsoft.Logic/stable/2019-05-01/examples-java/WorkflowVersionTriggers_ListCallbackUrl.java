import com.azure.resourcemanager.logic.models.GetCallbackUrlParameters;
import com.azure.resourcemanager.logic.models.KeyType;
import java.time.OffsetDateTime;

/** Samples for WorkflowVersionTriggers ListCallbackUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowVersionTriggers_ListCallbackUrl.json
     */
    /**
     * Sample code: Get the callback url for a trigger of a workflow version.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getTheCallbackUrlForATriggerOfAWorkflowVersion(
        com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowVersionTriggers()
            .listCallbackUrlWithResponse(
                "testResourceGroup",
                "testWorkflowName",
                "testWorkflowVersionId",
                "testTriggerName",
                new GetCallbackUrlParameters()
                    .withNotAfter(OffsetDateTime.parse("2017-03-05T08:00:00Z"))
                    .withKeyType(KeyType.PRIMARY),
                com.azure.core.util.Context.NONE);
    }
}
