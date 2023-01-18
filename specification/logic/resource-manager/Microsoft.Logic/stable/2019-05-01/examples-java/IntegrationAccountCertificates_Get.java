/** Samples for IntegrationAccountCertificates Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountCertificates_Get.json
     */
    /**
     * Sample code: Get certificate by name.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getCertificateByName(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountCertificates()
            .getWithResponse(
                "testResourceGroup", "testIntegrationAccount", "testCertificate", com.azure.core.util.Context.NONE);
    }
}
