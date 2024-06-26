
import com.azure.resourcemanager.informaticadatamanagement.models.AdvancedCustomProperties;
import com.azure.resourcemanager.informaticadatamanagement.models.ApplicationConfigs;
import com.azure.resourcemanager.informaticadatamanagement.models.ApplicationType;
import com.azure.resourcemanager.informaticadatamanagement.models.CdiConfigProps;
import com.azure.resourcemanager.informaticadatamanagement.models.InformaticaServerlessRuntimeProperties;
import com.azure.resourcemanager.informaticadatamanagement.models.NetworkInterfaceConfiguration;
import com.azure.resourcemanager.informaticadatamanagement.models.PlatformType;
import com.azure.resourcemanager.informaticadatamanagement.models.ServerlessRuntimeConfigProperties;
import com.azure.resourcemanager.informaticadatamanagement.models.ServerlessRuntimeNetworkProfile;
import com.azure.resourcemanager.informaticadatamanagement.models.ServerlessRuntimeTag;
import com.azure.resourcemanager.informaticadatamanagement.models.ServerlessRuntimeUserContextProperties;
import java.util.Arrays;

/**
 * Samples for ServerlessRuntimes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/informatica/Informatica.DataManagement/examples/2024-05-08/
     * ServerlessRuntimes_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: ServerlessRuntimes_CreateOrUpdate.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void serverlessRuntimesCreateOrUpdate(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager
            .serverlessRuntimes().define(
                "0j-__")
            .withExistingOrganization("rgopenapi",
                "__C")
            .withProperties(
                new InformaticaServerlessRuntimeProperties().withDescription("mqkaenjmxakvzrwmirelmhgiedto")
                    .withPlatform(PlatformType.AZURE).withApplicationType(ApplicationType.CDI)
                    .withComputeUnits("bsctukmndvowse").withExecutionTimeout("ruiougpypny")
                    .withServerlessAccountLocation("bkxdfopapbqucyhduewrubjpaei")
                    .withServerlessRuntimeNetworkProfile(
                        new ServerlessRuntimeNetworkProfile()
                            .withNetworkInterfaceConfiguration(
                                new NetworkInterfaceConfiguration()
                                    .withVnetId("uaqjvtubxccjs").withSubnetId("s").withVnetResourceGuid(
                                        "5328d299-1462-4be0-bef1-303a28e556a0")))
                    .withAdvancedCustomProperties(Arrays.asList(
                        new AdvancedCustomProperties().withKey("fakeTokenPlaceholder").withValue("unraxmnohdmvutt")))
                    .withSupplementaryFileLocation("zmlqtkncwgqhhupsnqluumz")
                    .withServerlessRuntimeConfig(new ServerlessRuntimeConfigProperties()
                        .withCdiConfigProps(Arrays.asList(new CdiConfigProps()
                            .withEngineName("hngsdqvtjdhwqlbqfotipaiwjuys").withEngineVersion("zlrlbg")
                            .withApplicationConfigs(Arrays.asList(new ApplicationConfigs()
                                .withType("lw").withName("upfvjrqcrwwedfujkmsodeinw").withValue("mozgsetpwjmtyl")
                                .withPlatform("dixfyeobngivyvf").withCustomized("j").withDefaultValue("zvgkqwmi")))))
                        .withCdieConfigProps(Arrays.asList(new CdiConfigProps()
                            .withEngineName("hngsdqvtjdhwqlbqfotipaiwjuys").withEngineVersion("zlrlbg")
                            .withApplicationConfigs(Arrays.asList(new ApplicationConfigs().withType("lw")
                                .withName("upfvjrqcrwwedfujkmsodeinw").withValue("mozgsetpwjmtyl")
                                .withPlatform("dixfyeobngivyvf").withCustomized("j").withDefaultValue("zvgkqwmi"))))))
                    .withServerlessRuntimeTags(
                        Arrays.asList(new ServerlessRuntimeTag().withName("korveuycuwhs").withValue("uyiuegxnkgp")))
                    .withServerlessRuntimeUserContextProperties(
                        new ServerlessRuntimeUserContextProperties().withUserContextToken("fakeTokenPlaceholder")))
            .create();
    }
}
