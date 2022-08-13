import com.azure.resourcemanager.automation.models.ContentHash;
import com.azure.resourcemanager.automation.models.ContentSource;
import com.azure.resourcemanager.automation.models.ContentSourceType;

/** Samples for DscConfiguration CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/createOrUpdateDscConfiguration.json
     */
    /**
     * Sample code: Create or Update Configuration.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createOrUpdateConfiguration(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscConfigurations()
            .define("SetupServer")
            .withExistingAutomationAccount("rg", "myAutomationAccount18")
            .withSource(
                new ContentSource()
                    .withHash(
                        new ContentHash()
                            .withAlgorithm("sha256")
                            .withValue("A9E5DB56BA21513F61E0B3868816FDC6D4DF5131F5617D7FF0D769674BD5072F"))
                    .withType(ContentSourceType.EMBEDDED_CONTENT)
                    .withValue(
                        "Configuration SetupServer {\r\n"
                            + "    Node localhost {\r\n"
                            + "                               WindowsFeature IIS {\r\n"
                            + "                               Name = \"Web-Server\";\r\n"
                            + "            Ensure = \"Present\"\r\n"
                            + "        }\r\n"
                            + "    }\r\n"
                            + "}"))
            .withRegion("East US 2")
            .withName("SetupServer")
            .withDescription("sample configuration")
            .create();
    }
}
