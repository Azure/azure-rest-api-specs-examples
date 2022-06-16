import java.time.Duration;

/** Samples for Users CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Users/putUser.json
     */
    /**
     * Sample code: putUser.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void putUser(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager
            .users()
            .define("testuser")
            .withExistingLab("testrg123", "testlab")
            .withEmail("testuser@contoso.com")
            .withAdditionalUsageQuota(Duration.parse("20:00"))
            .create();
    }
}
