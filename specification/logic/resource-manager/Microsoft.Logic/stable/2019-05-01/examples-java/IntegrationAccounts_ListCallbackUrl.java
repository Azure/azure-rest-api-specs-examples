import com.azure.resourcemanager.logic.models.GetCallbackUrlParameters;
import com.azure.resourcemanager.logic.models.KeyType;
import java.time.OffsetDateTime;

/** Samples for IntegrationAccounts ListCallbackUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccounts_ListCallbackUrl.json
     */
    /**
     * Sample code: List IntegrationAccount callback URL.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listIntegrationAccountCallbackURL(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccounts()
            .listCallbackUrlWithResponse(
                "testResourceGroup",
                "testIntegrationAccount",
                new GetCallbackUrlParameters()
                    .withNotAfter(OffsetDateTime.parse("2017-03-05T08:00:00Z"))
                    .withKeyType(KeyType.PRIMARY),
                com.azure.core.util.Context.NONE);
    }
}
