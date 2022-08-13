import com.azure.resourcemanager.automation.models.ContentHash;
import com.azure.resourcemanager.automation.models.ContentSource;
import com.azure.resourcemanager.automation.models.ContentSourceType;
import com.azure.resourcemanager.automation.models.DscConfigurationAssociationProperty;

/** Samples for DscNodeConfiguration CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateDscNodeConfiguration.json
     */
    /**
     * Sample code: Create node configuration.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createNodeConfiguration(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscNodeConfigurations()
            .define("configName.nodeConfigName")
            .withExistingAutomationAccount("rg", "myAutomationAccount20")
            .withName("configName.nodeConfigName")
            .withSource(
                new ContentSource()
                    .withHash(
                        new ContentHash()
                            .withAlgorithm("sha256")
                            .withValue("6DE256A57F01BFA29B88696D5E77A383D6E61484C7686E8DB955FA10ACE9FFE5"))
                    .withType(ContentSourceType.EMBEDDED_CONTENT)
                    .withValue(
                        "\r\n"
                            + "instance of MSFT_RoleResource as $MSFT_RoleResource1ref\r\n"
                            + "{\r\n"
                            + "ResourceID = \"[WindowsFeature]IIS\";\r\n"
                            + " Ensure = \"Present\";\r\n"
                            + " SourceInfo = \"::3::32::WindowsFeature\";\r\n"
                            + " Name = \"Web-Server\";\r\n"
                            + " ModuleName = \"PsDesiredStateConfiguration\";\r\n\r\n"
                            + "ModuleVersion = \"1.0\";\r\r\n"
                            + " ConfigurationName = \"configName\";\r\r\n"
                            + "};\r\n"
                            + "instance of OMI_ConfigurationDocument\r\n\r\r\n"
                            + "                    {\r\n"
                            + " Version=\"2.0.0\";\r\n"
                            + " \r\r\n"
                            + "                        MinimumCompatibleVersion = \"1.0.0\";\r\n"
                            + " \r\r\n"
                            + "                        CompatibleVersionAdditionalProperties="
                            + " {\"Omi_BaseResource:ConfigurationName\"};\r\n"
                            + " \r\r\n"
                            + "                        Author=\"weijiel\";\r\n"
                            + " \r\r\n"
                            + "                        GenerationDate=\"03/30/2017 13:40:25\";\r\n"
                            + " \r\r\n"
                            + "                        GenerationHost=\"TEST-BACKEND\";\r\n"
                            + " \r\r\n"
                            + "                        Name=\"configName\";\r\n\r\r\n"
                            + "                    };\r\n")
                    .withVersion("1.0"))
            .withConfiguration(new DscConfigurationAssociationProperty().withName("configName"))
            .withIncrementNodeConfigurationBuild(true)
            .create();
    }
}
