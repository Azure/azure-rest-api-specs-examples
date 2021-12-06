Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-desktopvirtualization_1.0.0-beta.1/sdk/desktopvirtualization/azure-resourcemanager-desktopvirtualization/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.desktopvirtualization.models.MsixPackageApplications;
import com.azure.resourcemanager.desktopvirtualization.models.MsixPackageDependencies;
import java.time.OffsetDateTime;
import java.util.Arrays;

/** Samples for MsixPackages CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/MsixPackage_Create.json
     */
    /**
     * Sample code: MSIXPackage_Create.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void mSIXPackageCreate(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .msixPackages()
            .define("msixpackagefullname")
            .withExistingHostPool("resourceGroup1", "hostpool1")
            .withImagePath("imagepath")
            .withPackageName("MsixPackage_name")
            .withPackageFamilyName("MsixPackage_FamilyName")
            .withDisplayName("displayname")
            .withPackageRelativePath("packagerelativepath")
            .withIsRegularRegistration(false)
            .withIsActive(false)
            .withPackageDependencies(
                Arrays
                    .asList(
                        new MsixPackageDependencies()
                            .withDependencyName("MsixTest_Dependency_Name")
                            .withPublisher("PublishedName")
                            .withMinVersion("version")))
            .withVersion("version")
            .withLastUpdated(OffsetDateTime.parse("2008-09-22T14:01:54.9571247Z"))
            .withPackageApplications(
                Arrays
                    .asList(
                        new MsixPackageApplications()
                            .withAppId("ApplicationId")
                            .withDescription("application-desc")
                            .withAppUserModelId("AppUserModelId")
                            .withFriendlyName("friendlyname")
                            .withIconImageName("Apptile")
                            .withRawIcon("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo".getBytes())
                            .withRawPng("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo".getBytes())))
            .create();
    }
}
```
