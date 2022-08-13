import com.azure.resourcemanager.automation.models.RunbookAssociationProperty;
import java.time.OffsetDateTime;

/** Samples for Webhook CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/examples/createOrUpdateWebhook.json
     */
    /**
     * Sample code: Create or update webhook.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createOrUpdateWebhook(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .webhooks()
            .define("TestWebhook")
            .withExistingAutomationAccount("rg", "myAutomationAccount33")
            .withName("TestWebhook")
            .withIsEnabled(true)
            .withUri("<uri>")
            .withExpiryTime(OffsetDateTime.parse("2018-03-29T22:18:13.7002872Z"))
            .withRunbook(new RunbookAssociationProperty().withName("TestRunbook"))
            .create();
    }
}
