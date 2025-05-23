
/**
 * Samples for IntegrationAccountCertificates List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/
     * IntegrationAccountCertificates_List.json
     */
    /**
     * Sample code: Get certificates by integration account name.
     * 
     * @param manager Entry point to LogicManager.
     */
    public static void getCertificatesByIntegrationAccountName(com.azure.resourcemanager.logic.LogicManager manager) {
        manager.integrationAccountCertificates().list("testResourceGroup", "testIntegrationAccount", null,
            com.azure.core.util.Context.NONE);
    }
}
