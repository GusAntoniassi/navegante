import * as React from 'react';
import { TiChevronLeft } from 'react-icons/ti';
import { AiOutlineSearch } from 'react-icons/ai';
import { Container as EntityContainer } from 'core';

import 'stylesheets/06-components/container-item.scss';

type Props = {
    container: EntityContainer | null;
}

const collapseContainerAttr = (e: React.MouseEvent) => {
    const target: Element = e.target as Element;
    const containerAttr = target.closest('.container-attribute');

    if (containerAttr?.classList.contains('-collapsed')) {
        containerAttr.classList.remove('-collapsed');
    } else {
        containerAttr?.classList.add('-collapsed');
    }
};

const Chevron: React.FC = () => {
    return <span className="chevron" onClick={collapseContainerAttr}><TiChevronLeft className="icon"/></span>;
};

export const Container: React.FC<Props> = ({
    container
}) => {
    if (container == null) return (null);

    return (
        <div className="container-item">
            <div className="heading">
                {container.id}
            </div>

            <ul className="content">
                <li className="container-attribute">Name: {container.name}</li>
                <li className="container-attribute -collapsed">
                    Ports <Chevron/>
                    <ul className="list">
                        {container.ports.map((port, index) => <li className="item" key={index}>- {port}</li>)}
                    </ul>
                </li>
                <li className="container-attribute -collapsed">
                    Volumes <Chevron/>
                    <ul className="list">
                        {container.volumes.map((volume, index) => <li className="item" key={index}>- {volume}</li>)}
                    </ul>
                </li>
                <li className="container-attribute -collapsed">
                    Networks <Chevron/>
                    <ul className="list">
                        {container.networks.map((network, index) => <li className="item" key={index}>- {network}</li>)}
                    </ul>
                </li>
                <li className="container-attribute -collapsed">
                    Stats <Chevron/>
                    <ul className="list">
                        <li className="item">CPU%: 0.01%</li>
                        <li className="item">Mem%: 0.15%</li>
                        <li className="item">Mem usg: 11.7MiB / 7.6GiB</li>
                        <li className="item">Net I/O: 12.9kB / 0B</li>
                        <li className="item">Block I/O: 38.1MB / 0B</li>
                    </ul>
                </li>
            </ul>

            <div className="footer text-right">
                <AiOutlineSearch className="cursor-pointer"/>
            </div>
        </div>
    );
};
