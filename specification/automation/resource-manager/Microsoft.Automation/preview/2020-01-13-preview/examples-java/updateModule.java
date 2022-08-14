import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.models.ContentHash;
import com.azure.resourcemanager.automation.models.ContentLink;
import com.azure.resourcemanager.automation.models.ModuleUpdateParameters;

/** Samples for Module Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateModule.json
     */
    /**
     * Sample code: Update a module.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void updateAModule(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .modules()
            .updateWithResponse(
                "rg",
                "MyAutomationAccount",
                "MyModule",
                new ModuleUpdateParameters()
                    .withContentLink(
                        new ContentLink()
                            .withUri("https://teststorage.blob.core.windows.net/mycontainer/MyModule.zip")
                            .withContentHash(
                                new ContentHash()
                                    .withAlgorithm("sha265")
                                    .withValue("07E108A962B81DD9C9BAA89BB47C0F6EE52B29E83758B07795E408D258B2B87A"))
                            .withVersion("1.0.0.0")),
                Context.NONE);
    }
}
