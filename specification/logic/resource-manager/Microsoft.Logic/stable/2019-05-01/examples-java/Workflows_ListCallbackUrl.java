import com.azure.resourcemanager.logic.models.GetCallbackUrlParameters;
import com.azure.resourcemanager.logic.models.KeyType;
import java.time.OffsetDateTime;

/** Samples for Workflows ListCallbackUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_ListCallbackUrl.json
     */
    /**
     * Sample code: Get callback url.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getCallbackUrl(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflows()
            .listCallbackUrlWithResponse(
                "testResourceGroup",
                "testWorkflow",
                new GetCallbackUrlParameters()
                    .withNotAfter(OffsetDateTime.parse("2018-04-19T16:00:00Z"))
                    .withKeyType(KeyType.PRIMARY),
                com.azure.core.util.Context.NONE);
    }
}
