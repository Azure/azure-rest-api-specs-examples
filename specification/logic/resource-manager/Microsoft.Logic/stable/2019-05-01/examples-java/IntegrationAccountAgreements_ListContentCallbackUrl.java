import com.azure.resourcemanager.logic.models.GetCallbackUrlParameters;
import com.azure.resourcemanager.logic.models.KeyType;
import java.time.OffsetDateTime;

/** Samples for IntegrationAccountAgreements ListContentCallbackUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAgreements_ListContentCallbackUrl.json
     */
    /**
     * Sample code: Get the content callback url.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getTheContentCallbackUrl(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountAgreements()
            .listContentCallbackUrlWithResponse(
                "testResourceGroup",
                "testIntegrationAccount",
                "testAgreement",
                new GetCallbackUrlParameters()
                    .withNotAfter(OffsetDateTime.parse("2018-04-19T16:00:00Z"))
                    .withKeyType(KeyType.PRIMARY),
                com.azure.core.util.Context.NONE);
    }
}
