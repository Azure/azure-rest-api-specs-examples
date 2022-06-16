const { ImageBuilderClient } = require("@azure/arm-imagebuilder");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a virtual machine image template
 *
 * @summary Create or update a virtual machine image template
 * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/CreateImageTemplateWindows.json
 */
async function createAnImageTemplateForWindows() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const imageTemplateName = "myImageTemplate";
  const parameters = {
    customize: [
      {
        name: "PowerShell (inline) Customizer Example",
        type: "PowerShell",
        inline: ["Powershell command-1", "Powershell command-2", "Powershell command-3"],
      },
      {
        name: "PowerShell (inline) Customizer Elevated user Example",
        type: "PowerShell",
        inline: ["Powershell command-1", "Powershell command-2", "Powershell command-3"],
        runElevated: true,
      },
      {
        name: "PowerShell (inline) Customizer Elevated Local System user Example",
        type: "PowerShell",
        inline: ["Powershell command-1", "Powershell command-2", "Powershell command-3"],
        runAsSystem: true,
        runElevated: true,
      },
      {
        name: "PowerShell (script) Customizer Example",
        type: "PowerShell",
        scriptUri: "https://example.com/path/to/script.ps1",
        validExitCodes: [0, 1],
      },
      {
        name: "PowerShell (script) Customizer Elevated Local System user Example",
        type: "PowerShell",
        runElevated: true,
        scriptUri: "https://example.com/path/to/script.ps1",
        validExitCodes: [0, 1],
      },
      {
        name: "PowerShell (script) Customizer Elevated Local System user Example",
        type: "PowerShell",
        runAsSystem: true,
        runElevated: true,
        scriptUri: "https://example.com/path/to/script.ps1",
        validExitCodes: [0, 1],
      },
      {
        name: "Restart Customizer Example",
        type: "WindowsRestart",
        restartCheckCommand: "powershell -command \"& {Write-Output 'restarted.'}\"",
        restartCommand: 'shutdown /f /r /t 0 /c "packer restart"',
        restartTimeout: "10m",
      },
      {
        name: "Windows Update Customizer Example",
        type: "WindowsUpdate",
        filters: ["$_.BrowseOnly"],
        searchCriteria: "BrowseOnly=0 and IsInstalled=0",
        updateLimit: 100,
      },
    ],
    distribute: [
      {
        type: "ManagedImage",
        artifactTags: { tagName: "value" },
        imageId:
          "/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Compute/images/image_it_1",
        location: "1_location",
        runOutputName: "image_it_pir_1",
      },
    ],
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000000000000000000000000000/resourcegroups/rg1/providers/MicrosoftManagedIdentity/userAssignedIdentities/identity1":
          {},
      },
    },
    location: "westus",
    source: {
      type: "ManagedImage",
      imageId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/images/source_image",
    },
    tags: { imagetemplateTag1: "IT_T1", imagetemplateTag2: "IT_T2" },
    vmProfile: {
      osDiskSizeGB: 64,
      vmSize: "Standard_D2s_v3",
      vnetConfig: {
        subnetId:
          "/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet_name/subnets/subnet_name",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ImageBuilderClient(credential, subscriptionId);
  const result = await client.virtualMachineImageTemplates.beginCreateOrUpdateAndWait(
    resourceGroupName,
    imageTemplateName,
    parameters
  );
  console.log(result);
}

createAnImageTemplateForWindows().catch(console.error);
