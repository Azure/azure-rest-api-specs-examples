```javascript
const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function iscsiDisksCreateOrUpdate() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-0NZI14MDTF";
  const iscsiServerName = "HSDK-0NZI14MDTF";
  const diskName = "Auto-TestIscsiDisk1";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const iscsiDisk = {
    name: "Auto-TestIscsiDisk1",
    description: "Demo IscsiDisk for SDK Test Tiered",
    accessControlRecords: [],
    dataPolicy: "Tiered",
    diskStatus: "Online",
    monitoringStatus: "Enabled",
    provisionedCapacityInBytes: 536870912000,
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.iscsiDisks.beginCreateOrUpdateAndWait(
    deviceName,
    iscsiServerName,
    diskName,
    resourceGroupName,
    managerName,
    iscsiDisk
  );
  console.log(result);
}

iscsiDisksCreateOrUpdate().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storsimple1200series_2.0.1/sdk/storsimple1200series/arm-storsimple1200series/README.md) on how to add the SDK to your project and authenticate.
