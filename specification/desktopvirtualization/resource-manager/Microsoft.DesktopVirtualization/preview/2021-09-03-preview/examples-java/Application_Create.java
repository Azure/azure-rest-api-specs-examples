import com.azure.resourcemanager.desktopvirtualization.models.CommandLineSetting;

/** Samples for Applications CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/Application_Create.json
     */
    /**
     * Sample code: Application_Create.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void applicationCreate(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .applications()
            .define("application1")
            .withExistingApplicationGroup("resourceGroup1", "applicationGroup1")
            .withCommandLineSetting(CommandLineSetting.ALLOW)
            .withDescription("des1")
            .withFriendlyName("friendly")
            .withFilePath("path")
            .withCommandLineArguments("arguments")
            .withShowInPortal(true)
            .withIconPath("icon")
            .withIconIndex(1)
            .create();
    }
}
