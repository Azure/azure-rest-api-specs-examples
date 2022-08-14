import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.models.ContentHash;
import com.azure.resourcemanager.automation.models.ContentLink;
import com.azure.resourcemanager.automation.models.ModuleCreateOrUpdateParameters;

/** Samples for Module CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateModule.json
     */
    /**
     * Sample code: Create or update a module.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createOrUpdateAModule(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .modules()
            .createOrUpdateWithResponse(
                "rg",
                "myAutomationAccount33",
                "OmsCompositeResources",
                new ModuleCreateOrUpdateParameters()
                    .withContentLink(
                        new ContentLink()
                            .withUri("https://teststorage.blob.core.windows.net/dsccomposite/OmsCompositeResources.zip")
                            .withContentHash(
                                new ContentHash()
                                    .withAlgorithm("sha265")
                                    .withValue("07E108A962B81DD9C9BAA89BB47C0F6EE52B29E83758B07795E408D258B2B87A"))
                            .withVersion("1.0.0.0")),
                Context.NONE);
    }
}
